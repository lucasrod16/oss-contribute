import React, { useEffect, useState } from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import Card from "./components/Card";
import Pagination from "./components/Pagination";
import Navbar from "./components/Navbar";
import Header from "./components/Header";
import Footer from "./components/Footer";
import Filter from "./components/Filter";
import { Language } from "./languages.ts";
import { Repo } from "./types/repo.ts";

const theme = createTheme({
	palette: {
		background: {
			default: "#0d1117",
		},
		text: {
			primary: "#c9d1d9",
			secondary: "#8b949e",
		},
		primary: {
			main: "#58a6ff",
		},
		secondary: {
			main: "#161b22",
		},
	},
	typography: {
		fontFamily: "Arial, sans-serif",
		h1: {
			fontSize: "2.5rem",
			fontWeight: 700,
		},
		body1: {
			fontSize: "1.15rem",
			lineHeight: 1.8,
		},
	},
});

const App = () => {
	const [cards, setCards] = useState<Repo[]>([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [selectedLanguages, setSelectedLanguages] = useState<Language[]>([]);
	const cardsPerPage = 20;

	useEffect(() => {
		const fetchRepos = async () => {
			const baseURL = window.location.hostname === "localhost" ? "http://localhost:8080" : "https://osscontribute.com";
			const apiURL = `${baseURL}/repos`;
			try {
				const response = await fetch(apiURL);
				if (!response.ok) {
					throw new Error(`HTTP error! Status: ${response.status}`);
				}
				const data = await response.json();
				setCards(data);
			} catch (error) {
				console.error("Error fetching repos:", (error as Error).message);
			}
		};
		fetchRepos();
	}, []);

	const filteredCards = selectedLanguages.length
		? cards.filter((repo) => selectedLanguages.includes(repo.language as Language))
		: cards;

	const indexOfLastCard = currentPage * cardsPerPage;
	const indexOfFirstCard = indexOfLastCard - cardsPerPage;
	const currentCards = filteredCards.slice(indexOfFirstCard, indexOfLastCard);

	const handleFilterChange = (selectedLanguages: Language[]) => {
		setSelectedLanguages(selectedLanguages);
		setCurrentPage(1); // Reset to first page when filter changes
	};

	return (
		<ThemeProvider theme={theme}>
			<CssBaseline />
			<Box sx={{ display: "flex", flexDirection: "column", minHeight: "100vh" }}>
				<Navbar />
				<Header />
				<Container sx={{ flex: 1 }}>
					<Filter selectedLanguages={selectedLanguages} onFilterChange={handleFilterChange} />
					<Box sx={{ display: "flex", flexDirection: "column", gap: 3, alignItems: "center" }}>
						{currentCards.map((repo) => (
							<Card repo={repo} key={repo.name} />
						))}
					</Box>
					<Pagination
						totalPages={Math.ceil(filteredCards.length / cardsPerPage)}
						currentPage={currentPage}
						onPageChange={(page: React.SetStateAction<number>) => setCurrentPage(page)}
					/>
				</Container>
				<Footer />
			</Box>
		</ThemeProvider>
	);
};

export default App;
