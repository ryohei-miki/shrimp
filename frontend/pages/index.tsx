import { useEffect, useState } from "react";

interface Evi {
  message: string;
}
export default function Home() {
  const [shrimp, setShrimp] = useState<string>("");
  const fetchEvi = async () => {
    const res = await fetch("http://localhost:3000", {});
    return await res.json();
  };
  useEffect(() => {
    fetchEvi().then((msg: Evi) => {
      setShrimp(msg.message);
    });
  }, []);
  return <div>{shrimp}</div>;
}
