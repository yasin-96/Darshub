// https://docs.cypress.io/api/introduction/api.html

describe("TEST: Registration View", () => {
  it("Visits Registration View", () => {
    cy.visit("/registry");
  });

  it("BaseData From should exist", () => {
    cy.visit("/registry");
    cy.get("form").should("exist");
  });

  it("Inputfields shoud have all inportant data filled out", () => {
    cy.visit("/registry");
    cy.get('[data-test="first_name"]')
      .type("Alex")
      .get('[data-test="last_name"]')
      .type("J.")
    .get('[data-test="next_baseData"]').should("disabled")
  });

  it("Clear should reset the form", () => {
    cy
      .visit("/registry")
      .get('[data-test="first_name"]')
      .type("Alex")
      .get('[data-test="last_name"]')
      .type("J.")
      .get('[data-test="abbort_baseData"]').click()
      .visit("/registry")
      .get('[data-test="first_name"]').should('be.empty')
      .get('[data-test="last_name"]').should('be.empty')
      // .get('[data-test="birthday"]').should('be.empty')
      .get('[data-test="bio"]').should('be.empty')
      .get('[data-test="country"]').should('be.empty')
      .get('[data-test="next_baseData"]').should("disabled")

  });
});
