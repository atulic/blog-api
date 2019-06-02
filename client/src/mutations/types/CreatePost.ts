/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: CreatePost
// ====================================================

export interface CreatePost_create {
  __typename: "Post";
  title: string | null;
  content: string | null;
}

export interface CreatePost {
  /**
   * Create new post
   */
  create: CreatePost_create | null;
}

export interface CreatePostVariables {
  title: string;
  content: string;
}
