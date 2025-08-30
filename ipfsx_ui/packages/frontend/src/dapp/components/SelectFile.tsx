
// 一个选择文件上传功能
import { useState } from 'react';

const SelectFilePage = () => {
  const [file, setFile] = useState<File | null>(null);
  const [fileName, setFileName] = useState('file');

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files.length > 0) {
      setFile(files[0]);
      setFileName(files[0].name);
    }

  }

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
      {file && <p>Selected file: {fileName}</p>}
    </div>
  )

}

export default SelectFilePage;