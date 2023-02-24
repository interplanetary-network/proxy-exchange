import Layout from "../Layout";
import PoolForm from "./PoolForm";

export default function PoolNew() {
  return (
    <Layout>
      <div className="pool-new">
        <h2>Publish Pool</h2>
        <PoolForm />
      </div>
    </Layout>
  )
}
