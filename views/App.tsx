import { useEffect, useState } from "react";

const App = () => {
	const [status, setStatus] = useState("checking");

	useEffect(() => {
		fetch("/api/health")
			.then((response) => {
				if (!response.ok) throw new Error("API request failed");
				return response.json();
			})
			.then((data) => setStatus(data.status === "ok" ? "online" : "offline"))
			.catch(() => setStatus("offline"));
	}, []);

	return (
		<main className="app">
			<h1>React in Go</h1>
			<p>Go Fiber API on Vercel Serverless Functions</p>
			<p className={`status status-${status}`}>API status: {status}</p>
		</main>
	);
};

export default App;
