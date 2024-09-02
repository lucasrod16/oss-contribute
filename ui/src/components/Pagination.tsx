import { Pagination as MuiPagination, Stack } from "@mui/material";
import React from "react";

interface PaginationProps {
	totalPages: number;
	currentPage: number;
	onPageChange: (page: number) => void;
}

const Pagination: React.FC<PaginationProps> = ({ totalPages, currentPage, onPageChange }) => {
	return (
		<Stack spacing={2} sx={{ my: 4, justifyContent: "center", alignItems: "center" }}>
			<MuiPagination
				count={totalPages}
				page={currentPage}
				onChange={(event, page) => onPageChange(page)}
				color="primary"
			/>
		</Stack>
	);
};

export default Pagination;
