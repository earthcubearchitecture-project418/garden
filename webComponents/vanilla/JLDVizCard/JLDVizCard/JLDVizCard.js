class JLDVizCard extends HTMLElement {
  constructor() {
    // If you define a constructor, always call super() first as it is required by the CE spec.
    super();

    // Setup a click listener on <jldviz-card>
    this.addEventListener('click', e => {
      this.toggleCard();
    });
  }

  toggleCard() {
    console.log("Element was clicked!");
  }
}

customElements.define('jldviz-card', JLDVizCard);
