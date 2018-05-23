import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';

/**
 * `p418-search`
 * Project 418 search element
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class P418Search extends PolymerElement {
  static get template() {
    return html`
      <style>
        :host {
          display: block;
        }
      </style>
      <h2>Hello [[prop1]]!</h2>
    `;
  }
  static get properties() {
    return {
      prop1: {
        type: String,
        value: 'p418-search',
      },
    };
  }
}

window.customElements.define('p418-search', P418Search);
