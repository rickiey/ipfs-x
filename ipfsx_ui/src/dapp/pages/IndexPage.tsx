import { FC } from 'react'
import SelectFile from '~~/dapp/components/SelectFile'
import Layout from '~~/components/layout/Layout'
import NetworkSupportChecker from '../../components/NetworkSupportChecker'


const IndexPage: FC = () => {
  return (
    <Layout>
      <NetworkSupportChecker />
      <div className="justify-content flex flex-grow flex-col items-center justify-center rounded-md p-3">
        <SelectFile />
      </div>
    </Layout>
  )
}

export default IndexPage
