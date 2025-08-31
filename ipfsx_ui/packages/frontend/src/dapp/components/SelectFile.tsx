import { Button, Card, Progress, Text } from "@radix-ui/themes";
import { Upload, FileIcon, CheckCircle } from "lucide-react";
import { useState, useCallback } from "react";
import toast from "react-hot-toast";
import { create } from "ipfs-http-client";
import Notification from "~~/components/Notification";

interface UploadResult {
  hash: string;
  name: string;
  size: number;
}

const SelectFile = () => {
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [uploadProgress, setUploadProgress] = useState(0);
  const [uploadResult, setUploadResult] = useState<UploadResult | null>(null);
  const [dragOver, setDragOver] = useState(false);

  const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
  };

  const uploadToIPFS = async (file: File): Promise<UploadResult> => {
    // 模拟上传进度
    const progressInterval = setInterval(() => {
      setUploadProgress((prev: number) => {
        if (prev >= 90) {
          clearInterval(progressInterval);
          return 90;
        }
        return prev + Math.random() * 20;
      });
    }, 200);

    try {
      // 创建 IPFS HTTP 客户端
      const ipfs = create({
        url: "http://localhost:5001",
        timeout: 60000, // 60秒超时
      });

      // 上传文件到 IPFS
      const result = await ipfs.add(file, {
        progress: (prog: number) => {
          // 更新真实的上传进度
          const percentage = Math.round((prog / file.size) * 100);
          setUploadProgress(Math.min(percentage, 90));
        },
      });

      clearInterval(progressInterval);
      setUploadProgress(100);

      return {
        hash: result.cid.toString(),
        name: file.name,
        size: file.size,
      };
    } catch (error) {
      clearInterval(progressInterval);
      console.error("IPFS upload error:", error);
      throw new Error(
        error instanceof Error
          ? `IPFS 上传失败: ${error.message}`
          : "IPFS 上传失败: 未知错误"
      );
    }
  };

  const handleFileSelect = useCallback((selectedFile: File) => {
    setFile(selectedFile);
    setUploadResult(null);
    setUploadProgress(0);
  }, []);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files.length > 0) {
      handleFileSelect(files[0]);
    }
  };

  const handleDrop = useCallback(
    (e: React.DragEvent<HTMLDivElement>) => {
      e.preventDefault();
      setDragOver(false);

      const files = e.dataTransfer.files;
      if (files && files.length > 0) {
        handleFileSelect(files[0]);
      }
    },
    [handleFileSelect]
  );

  const handleDragOver = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragOver(true);
  }, []);

  const handleDragLeave = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragOver(false);
  }, []);

  const handleUpload = async () => {
    if (!file) return;

    setUploading(true);
    setUploadProgress(0);

    const toastId = toast.custom(
      <Notification type="loading">正在上传到 IPFS...</Notification>
    );

    try {
      const result = await uploadToIPFS(file);
      setUploadResult(result);

      toast.dismiss(toastId);
      toast.custom(
        <Notification type="success">
          文件上传成功！IPFS Hash: {result.hash.substring(0, 20)}...
        </Notification>
      );
    } catch (error) {
      console.error("Upload error:", error);
      toast.dismiss(toastId);
      toast.custom(
        <Notification type="error">
          上传失败: {error instanceof Error ? error.message : "未知错误"}
        </Notification>
      );
    } finally {
      setUploading(false);
    }
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
    toast.custom(<Notification type="success">已复制到剪贴板</Notification>);
  };

  return (
    <div className="w-full max-w-2xl space-y-6">
      {/* 文件选择区域 */}
      <Card className="p-6">
        <div
          className={`border-2 border-dashed rounded-lg p-8 text-center transition-colors ${
            dragOver
              ? "border-blue-400 bg-blue-50 dark:bg-blue-950"
              : "border-gray-300 dark:border-gray-600"
          }`}
          onDrop={handleDrop}
          onDragOver={handleDragOver}
          onDragLeave={handleDragLeave}
        >
          <Upload className="mx-auto h-12 w-12 text-gray-400 mb-4" />
          <Text size="4" className="block mb-2">
            拖拽文件到此处或点击选择文件
          </Text>
          <input
            type="file"
            onChange={handleFileChange}
            className="hidden"
            id="file-input"
          />
          <Button asChild variant="outline">
            <label htmlFor="file-input" className="cursor-pointer">
              选择文件
            </label>
          </Button>
        </div>
      </Card>

      {/* 文件信息 */}
      {file && (
        <Card className="p-6">
          <div className="flex items-center space-x-4">
            <FileIcon className="h-8 w-8 text-blue-500" />
            <div className="flex-1">
              <Text size="3" weight="medium" className="block">
                {file.name}
              </Text>
              <Text size="2" color="gray">
                {formatFileSize(file.size)}
              </Text>
            </div>
            <Button onClick={handleUpload} disabled={uploading} size="3">
              {uploading ? "上传中..." : "上传到 IPFS"}
            </Button>
          </div>

          {/* 上传进度 */}
          {uploading && (
            <div className="mt-4">
              <Progress value={uploadProgress} className="w-full" />
              <Text size="2" color="gray" className="mt-1">
                {Math.round(uploadProgress)}% 完成
              </Text>
            </div>
          )}
        </Card>
      )}

      {/* 上传结果 */}
      {uploadResult && (
        <Card className="p-6">
          <div className="flex items-start space-x-3">
            <CheckCircle className="h-6 w-6 text-green-500 mt-1" />
            <div className="flex-1 space-y-3">
              <Text size="3" weight="medium" className="text-green-600">
                上传成功！
              </Text>

              <div className="space-y-2">
                <div>
                  <Text size="2" color="gray">
                    IPFS Hash:
                  </Text>
                  <div className="flex items-center space-x-2 mt-1">
                    <code className="bg-gray-100 dark:bg-gray-800 px-2 py-1 rounded text-sm">
                      {uploadResult.hash}
                    </code>
                    <Button
                      size="1"
                      variant="ghost"
                      onClick={() => copyToClipboard(uploadResult.hash)}
                    >
                      复制
                    </Button>
                  </div>
                </div>

                <div>
                  <Text size="2" color="gray">
                    IPFS Gateway URL:
                  </Text>
                  <div className="flex items-center space-x-2 mt-1">
                    <code className="bg-gray-100 dark:bg-gray-800 px-2 py-1 rounded text-sm break-all">
                      https://ipfs.io/ipfs/{uploadResult.hash}
                    </code>
                    <Button
                      size="1"
                      variant="ghost"
                      onClick={() =>
                        copyToClipboard(
                          `https://ipfs.io/ipfs/${uploadResult.hash}`
                        )
                      }
                    >
                      复制
                    </Button>
                  </div>
                </div>
              </div>

              <Button
                variant="outline"
                onClick={() =>
                  window.open(
                    `https://ipfs.io/ipfs/${uploadResult.hash}`,
                    "_blank"
                  )
                }
              >
                在 IPFS 中查看
              </Button>
            </div>
          </div>
        </Card>
      )}
    </div>
  );
};

export default SelectFile;
