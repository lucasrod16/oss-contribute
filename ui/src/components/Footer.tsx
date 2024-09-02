import React from "react";
import { Box, Typography, useTheme } from "@mui/material";

const Footer: React.FC = () => {
	const theme = useTheme();
	const currentYear = new Date().getFullYear();

	return (
		<Box
			sx={{
				p: 2,
				mt: "auto",
				backgroundColor: theme.palette.secondary.main,
				color: theme.palette.text.secondary,
				textAlign: "center",
				borderTop: `1px solid ${theme.palette.text.primary}`,
			}}
		>
			<Typography variant="body2">&copy; {currentYear} By Lucas Rodriguez. All rights reserved.</Typography>
		</Box>
	);
};

export default Footer;
