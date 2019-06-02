import React from "react";
import { Query } from "react-apollo";
import { POST_QUERY } from "../queries/fetchPostQuery";
import { FetchPosts } from "../queries/types/FetchPosts";
import { Loading } from "./Loading";
import { Error } from "./Error";
import { Box } from "@material-ui/core";

export const BlogPost: React.FC = () => (
  <Query<FetchPosts> query={POST_QUERY}>
    {({ loading, error, data }) => {
      if (loading) return <Loading />;
      if (error) return <Error />;
      return (
        data &&
        data.posts &&
        data.posts.map((post, i) => (
          <Box key={i}>
            <div>{post && post.title}</div>
            <div>{post && post.content}</div>
          </Box>
        ))
      );
    }}
  </Query>
);
