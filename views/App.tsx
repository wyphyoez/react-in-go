import React, { useEffect, useState } from "react";

type V2Response = {
	name: string;
	version: string;
	framework: string;
	status: string;
	timestamp: string;
	features: string[];
};

const App = () => {
	const [data, setData] = useState<V2Response | null>(null);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		const load = async () => {
			try {
				setLoading(true);
				const response = await fetch("/api/v2/message");
				if (!response.ok) {
					throw new Error(`Request failed: ${response.status}`);
				}
				const json: V2Response = await response.json();
				setData(json);
			} catch (err) {
				setError(err instanceof Error ? err.message : "Unknown error");
			} finally {
				setLoading(false);
			}
		};

		void load();
	}, []);

	return (
		<main className="page">
			<section className="card">
				<h1>React in Go — Version 2</h1>
				<p className="subtitle">Frontend + Backend integration demo</p>

				{loading && <p>Loading API response...</p>}
				{error && <p className="error">Error: {error}</p>}

				{data && (
					<>
						<div className="meta">
							<div>
								<strong>Project:</strong> {data.name}
							</div>
							<div>
								<strong>Version:</strong> {data.version}
							</div>
							<div>
								<strong>Stack:</strong> {data.framework}
							</div>
							<div>
								<strong>Status:</strong> <span className="ok">{data.status}</span>
							</div>
							<div>
								<strong>Timestamp:</strong> {new Date(data.timestamp).toLocaleString()}
							</div>
						</div>

						<h2>Features in v2</h2>
						<ul>
							{data.features.map((feature) => (
								<li key={feature}>{feature}</li>
							))}
						</ul>
					</>
				)}
			</section>
		</main>
	);
};

export default App;
