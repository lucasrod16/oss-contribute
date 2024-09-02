import { Language } from "../languages.ts";

export type Repo = {
	name: string;
	description: string;
	owner: string;
	repoURL: string;
	avatarURL: string;
	language: Language;
	stars: number;
};
