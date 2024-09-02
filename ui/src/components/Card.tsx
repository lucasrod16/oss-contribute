import React from "react";
import { Card as MuiCard, CardContent, Typography, CardMedia, Box } from "@mui/material";
import { getLanguageImageURL } from "../languages";
import { Repo } from "../types/repo";

const Card: React.FC<{ repo: Repo }> = ({ repo }) => {
	return (
		<MuiCard
			sx={{
				width: "100%",
				maxWidth: "90%",
				margin: "auto",
				borderRadius: 2,
				bgcolor: "secondary.main",
				boxShadow: 3,
				textDecoration: "none",
				overflow: "hidden",
				position: "relative",
				transition: "transform 0.2s ease-in-out",
				"&:hover": { transform: "translateY(-5px)" },
			}}
			component="a"
			href={repo.repoURL}
			target="_blank"
			rel="noopener noreferrer"
		>
			<img
				src={`https://img.shields.io/github/stars/${repo.owner}/${repo.name}.svg?style=social`}
				alt="stars"
				width="80"
				style={{
					position: "absolute",
					top: "8px",
					right: "8px",
					zIndex: 10,
				}}
			/>
			<img
				src={`https://img.shields.io/github/contributors-anon/${repo.owner}/${repo.name}.svg`}
				alt="contributors"
				width="100"
				style={{
					position: "absolute",
					bottom: "8px",
					right: "8px",
					zIndex: 10,
				}}
			/>

			<Box
				sx={{
					display: "flex",
					alignItems: "flex-start",
					position: "relative",
					paddingRight: { xs: 2, sm: 2, md: "150px" },
					paddingTop: 2,
				}}
			>
				<CardMedia
					component="img"
					image={repo.avatarURL}
					alt={repo.name}
					sx={{
						objectFit: "contain",
						width: { xs: "50px", sm: "60px" },
						height: { xs: "50px", sm: "60px" },
						margin: "8px",
						borderRadius: "50%",
					}}
				/>
				<CardContent
					sx={{
						flex: 1,
						padding: 2,
						paddingRight: { xs: 2, sm: 2, md: "150px" },
						paddingBottom: { xs: "50px", sm: "60px" },
					}}
				>
					<Typography
						variant="h6"
						component="div"
						sx={{
							mb: 1,
							fontSize: { xs: "1rem", sm: "1.25rem" },
							overflowWrap: "break-word",
							whiteSpace: "normal",
							textOverflow: "ellipsis",
						}}
					>
						{repo.name}
					</Typography>
					<Typography
						variant="body2"
						color="text.secondary"
						sx={{
							mb: 1,
							maxWidth: { xs: "100%", sm: "calc(100% - 120px)" },
							whiteSpace: "normal",
							wordBreak: "break-word",
							lineHeight: 1.5,
						}}
					>
						{repo.description}
					</Typography>
				</CardContent>
				<Box
					sx={{
						position: "absolute",
						bottom: "8px",
						left: "50%",
						transform: "translateX(-50%)",
						display: "flex",
						justifyContent: "center",
						width: "100%",
						p: 1,
					}}
				>
					<img
						src={getLanguageImageURL(repo.language)}
						alt={repo.language}
						style={{
							width: "20%",
							maxWidth: "35px",
							height: "auto",
						}}
					/>
				</Box>
			</Box>
		</MuiCard>
	);
};

export default Card;
