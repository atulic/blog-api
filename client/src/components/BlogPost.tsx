import gql from "graphql-tag";
import React from "react";
import { Query } from "react-apollo";

interface Data {
  posts: {
    title: string;
    content: string;
  };
}

const POST_QUERY = gql`
  {
    posts(id: 1) {
      title
      content
    }
  }
`;

export const BlogPost: React.FC = props => (
  <Query<Data> query={POST_QUERY}>
    {({ loading, error, data }) => {
      if (loading) return <div>Fetching...</div>;
      if (error) return <div>Errored</div>;
      return (
        <>
          <div>{data && data.posts.title}</div>
          <div>{data && data.posts.content}</div>
        </>
      );
    }}
  </Query>
);
