import { InMemoryCache } from "apollo-cache-inmemory";
import { ApolloClient } from "apollo-client";
import { createHttpLink } from "apollo-link-http";
import React from "react";
import { ApolloProvider } from "react-apollo";
import ReactDOM from "react-dom";
import * as serviceWorker from "./serviceWorker";
import "./styles/index.css";

const httpLink = createHttpLink({
    uri: "http://localhost:4000/graphql",
  });

const client = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache(),
  });

const App = () => (
    <ApolloProvider client={client}>
      <div>
        <h2>My first Apollo app <span role="img" aria-label="rocket">🚀</span></h2>
      </div>
    </ApolloProvider>
  );

ReactDOM.render(<App />, document.getElementById("root"));

serviceWorker.unregister();
