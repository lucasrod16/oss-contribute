import React from "react";
import { AppBar, Toolbar, Typography, useTheme } from "@mui/material";
import { alpha } from "@mui/material/styles";
import { Chip, Box  } from "@mui/material";
import GitHubIcon from '@mui/icons-material/GitHub';

const Navbar: React.FunctionComponent = () => {
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
			<Toolbar sx={{ display: "flex", justifyContent: "center" }} >
				<Box sx={{ marginRight: "auto", visibility: "hidden" }}></Box>
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
				<Chip variant="outlined" label="Source" component="a" href="https://github.com/lucasrod16/oss-contribute" icon={<GitHubIcon/>} clickable 
					sx={{ marginLeft: "auto" }} 
					/>
			</Toolbar>
		</AppBar>
	);
};

export default Navbar;
