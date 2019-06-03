import React from "react";
import { Query } from "react-apollo";
import { POST_QUERY } from "../../queries/fetchPostQuery";
import { FetchPosts, FetchPosts_posts } from "../../queries/types/FetchPosts";
import { Loading } from "../Loading";
import { Error } from "../Error";
import { Grid } from "@material-ui/core";
import { BlogCard } from "./BlogCard";

export const BlogPostList: React.FC = () => {
  return (
    <Query<FetchPosts> query={POST_QUERY}>
      {({ loading, error, data }) => {
        if (loading) return <Loading />;
        if (error) return <Error />;

        const postContent =
          data &&
          data.posts &&
          data.posts.map((post, i) => (
            <BlogCard id={i} post={post as FetchPosts_posts} />
          ));

        return (
          <Grid container spacing={2}>
            {postContent}
          </Grid>
        );
      }}
    </Query>
  );
};
