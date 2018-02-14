import { Component, Prop } from '@stencil/core';

@Component({
  tag: 'test-one',
  styleUrl: 'test-one.scss'
})
export class MyComponent {

  // Indicate that name should be a public property on the component
  @Prop() name: string;

  render() {
    return (
      <p>
        This is a test {this.name}
      </p>
    );
  }
}