import {InMemoryCache} from "apollo-cache-inmemory";
import {ApolloClient} from "apollo-client";
import {createHttpLink} from "apollo-link-http";
import React from "react";
import {ApolloProvider} from "react-apollo";
import ReactDOM from "react-dom";
import App from "./components/App";
import * as serviceWorker from "./serviceWorker";
import {CssBaseline} from "@material-ui/core";

const httpLink = createHttpLink({
    uri: "http://localhost:4000/graphql"
});

const client = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache()
});

const WrappedApp = () => (
    <ApolloProvider client={client}>
        <CssBaseline/>
        <App/>
    </ApolloProvider>
);

ReactDOM.render(<WrappedApp/>, document.getElementById("root"));

serviceWorker.unregister();
