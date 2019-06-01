import gql from "graphql-tag";
import React from "react";
import {Query} from "react-apollo";
import {Post} from "./types";

interface Data {
  posts: Post[];
}

const POST_QUERY = gql`
  {
    posts {
      title
      content
    }
  }
`;

export const BlogPost: React.FC = () => (
  <Query<Data> query={POST_QUERY}>
    {({ loading, error, data }) => {
      if (loading) return <div>Fetching...</div>;
      if (error) return <div>Errored</div>;
      return (
          data &&
          data.posts.map(post => (
              <>
                <div>{post.title}</div>
                <div>{post.content}</div>
              </>
          ))
      );
    }}
  </Query>
);
