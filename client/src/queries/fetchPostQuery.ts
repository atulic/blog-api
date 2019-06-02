import {gql} from "apollo-boost";

export const POST_QUERY = gql`
  query FetchPosts {
    posts {
      title
      content
    }
  }
`;