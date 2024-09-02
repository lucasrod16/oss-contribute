import React from "react";
import { AppBar, Toolbar, Typography, useTheme } from "@mui/material";
import { alpha } from "@mui/material/styles";

const Navbar: React.FC = () => {
	const theme = useTheme();

	const scrollToTop = () => {
		window.scrollTo({ top: 0, behavior: "smooth" });
	};

	return (
		<AppBar
			position="sticky"
			sx={{
				backdropFilter: "blur(10px)",
				backgroundColor: alpha(theme.palette.background.default, 0.7),
				boxShadow: "none",
				borderBottom: `1px solid ${alpha(theme.palette.text.primary, 0.12)}`,
				py: 0.5,
			}}
		>
			<Toolbar sx={{ display: "flex", justifyContent: "center" }}>
				<Typography
					variant="h6"
					onClick={scrollToTop}
					sx={{
						fontWeight: 700,
						color: theme.palette.text.primary,
						textTransform: "lowercase",
						letterSpacing: "0.1em",
						cursor: "pointer",
					}}
				>
					osscontribute.com
				</Typography>
			</Toolbar>
		</AppBar>
	);
};

export default Navbar;
