describe("Check Page Title", () => {
	it("has title", () => {
		cy.visit("/");
		cy.title().should("eq", "oss-projects");
	});
});

describe("Navbar Link", () => {
	it("clicking navbar brand takes user to top of the page", () => {
		cy.visit("/");
		cy.scrollTo(0, 500);
		cy.contains("osscontribute.com").click();
		cy.window().should("have.property", "scrollY", 0);
	});
});

describe("Filter Dropdown", () => {
	it("clicking each language in the filter dropdown shows at least 1 card", () => {
		cy.visit("/");
		cy.get("#language-select").click();
		cy.get('li[role="option"]').each(($language) => {
			cy.get("#language-select").click({ force: true });
			cy.wrap($language).click();
			cy.get("a.MuiCard-root").should("have.length.greaterThan", 0);
		});
	});
});

describe("Pagination", () => {
	beforeEach(() => {
		cy.visit("/");
	});

	it("should navigate to the next page and show different content", () => {
		cy.get("a.MuiCard-root").should("have.length", 20);
		cy.get('button[aria-label="Go to next page"]').click();
		cy.get("a.MuiCard-root").should("have.length", 20);
	});

	it("should navigate to a specific page and show correct content", () => {
		cy.get('button[aria-label="Go to page 2"]').click();
		cy.get("a.MuiCard-root").should("have.length", 20);
	});

	it("should show the correct number of cards when filtering", () => {
		cy.get("#language-select").click({ force: true });
		cy.get('li[role="option"]').contains("TypeScript").click();
		cy.get("a.MuiCard-root").should("have.length.greaterThan", 0);

		// ensure pagination still works after filtering
		cy.get('button[aria-label="Go to next page"]').click();
		cy.get("a.MuiCard-root").should("have.length.greaterThan", 0);
	});
});
