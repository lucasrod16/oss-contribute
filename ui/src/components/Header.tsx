import React from "react";
import { Container, Typography, Box, Link } from "@mui/material";
import { Fade } from "@mui/material";

const Header = () => {
	return (
		<Box sx={{ textAlign: "center", mt: 4, mb: 4, position: "relative" }}>
			<Box
				sx={{
					position: "absolute",
					top: 0,
					left: 0,
					width: "100%",
					height: "100%",
					backgroundColor: "background.default",
					zIndex: -1,
				}}
			/>
			<Container>
				<Fade in={true} timeout={1000}>
					<Box>
						<Typography variant="h1">Explore Top Open Source Projects</Typography>
						<Typography variant="body1" sx={{ mt: 2 }}>
							Contributing to open source is a great way to grow your skills and connect with talented people worldwide.
						</Typography>
						<Typography variant="body1">
							💡 Tip: Look for issues labeled "help wanted" and "good first issue" to find things to work on.
						</Typography>
						<Link
							href="https://docs.github.com/en/issues/tracking-your-work-with-issues/filtering-and-searching-issues-and-pull-requests#filtering-issues-and-pull-requests-by-labels"
							target="_blank"
							rel="noopener noreferrer"
							color="primary"
							underline="hover"
						>
							Learn more in the GitHub documentation
						</Link>
					</Box>
				</Fade>
			</Container>
		</Box>
	);
};

export default Header;
