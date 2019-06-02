import { gql } from "apollo-boost";

export const POST_MUTATION = gql`
  mutation CreatePost($title: String!, $content: String!) {
    create(title: $title, content: $content) {
      title
      content
    }
  }
`;
