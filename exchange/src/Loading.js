import { useEffect, useState } from "react";

export default function Loading({ withoutText = false}) {
  const [count, setCount] = useState(0);

  useEffect(() => {
    const intervalId = setInterval(() => {
      setCount(c => c < 3 ? (c + 1) : 1);
    }, 500);
    return () => clearInterval(intervalId);
  }, []);

  return <div className="loading">{withoutText ? "" : "Loading "}{[...Array(count).keys()].map(() => ".")}</div>
}
