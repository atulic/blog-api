/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: FetchPosts
// ====================================================

export interface FetchPosts_posts {
  __typename: "Post";
  title: string | null;
  content: string | null;
}

export interface FetchPosts {
  posts: (FetchPosts_posts | null)[] | null;
}
