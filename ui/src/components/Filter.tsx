import React, { useState } from "react";
import { Box, FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from "@mui/material";
import { Language, languageImageMap } from "../languages.ts";

type FilterProps = {
	selectedLanguages: Language[];
	onFilterChange: (selectedLanguages: Language[]) => void;
};

const Filter: React.FC<FilterProps> = ({ selectedLanguages, onFilterChange }) => {
	const [selectedLanguage, setSelectedLanguage] = useState<Language | "">("");

	const handleDropdownChange = (event: SelectChangeEvent<Language | "">) => {
		const selected = event.target.value as Language | "";

		if (selected === "") {
			setSelectedLanguage("");
			onFilterChange([]);
		} else {
			setSelectedLanguage(selected);
			onFilterChange([selected]);
		}
	};

	return (
		<Box sx={{ width: 200, mb: 4 }}>
			<FormControl fullWidth variant="outlined">
				<InputLabel id="language-select-label">Language</InputLabel>
				<Select
					labelId="language-select-label"
					id="language-select"
					value={selectedLanguage}
					label="Language"
					onChange={handleDropdownChange}
					sx={{
						backgroundColor: "secondary.main",
						color: "text.primary",
						borderRadius: 1,
						"& .MuiOutlinedInput-notchedOutline": {
							borderColor: "transparent",
						},
						"&:hover .MuiOutlinedInput-notchedOutline": {
							borderColor: "primary.main",
						},
						"&.Mui-focused .MuiOutlinedInput-notchedOutline": {
							borderColor: "primary.main",
						},
					}}
					MenuProps={{
						PaperProps: {
							sx: {
								backgroundColor: "#2d333b",
								color: "text.primary",
								"& .MuiMenuItem-root": {
									"&:hover": {
										backgroundColor: "#444c56",
									},
									"&.Mui-selected": {
										backgroundColor: "#444c56",
										"&:hover": {
											backgroundColor: "#586069",
										},
									},
								},
							},
						},
					}}
				>
					<MenuItem value="">All Languages</MenuItem>
					{Object.keys(languageImageMap).map((language) => (
						<MenuItem key={language} value={language}>
							{language}
						</MenuItem>
					))}
				</Select>
			</FormControl>
		</Box>
	);
};

export default Filter;
