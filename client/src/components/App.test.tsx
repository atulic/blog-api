import React from "react";
import { MockedProvider } from "react-apollo/test-utils";
import ReactDOM from "react-dom";
import App from "./App";

it("renders without crashing", () => {
  const div = document.createElement("div");

  const AppWithMockedProvider = () => (
    <MockedProvider>
      <App />
    </MockedProvider>
  );

  ReactDOM.render(<AppWithMockedProvider />, div);
  ReactDOM.unmountComponentAtNode(div);
});
