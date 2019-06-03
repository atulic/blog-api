import React from "react";
import { Query } from "react-apollo";
import { POST_QUERY } from "../queries/fetchPostQuery";
import { FetchPosts } from "../queries/types/FetchPosts";
import { Loading } from "./Loading";
import { Error } from "./Error";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Grid,
  Typography
} from "@material-ui/core";

export const BlogPost: React.FC = () => {
  return (
    <Query<FetchPosts> query={POST_QUERY}>
      {({ loading, error, data }) => {
        if (loading) return <Loading />;
        if (error) return <Error />;

        const postContent =
          data &&
          data.posts &&
          data.posts.map((post, i) => (
            <Grid item xs={12} sm={6} md={4}>
              <Card key={i}>
                <CardContent>
                  <Typography variant="h5">{post && post.title}</Typography>
                  <Typography variant="body1">
                    {post && post.content}
                  </Typography>
                </CardContent>
                <CardActions>
                  <Button size="small" color="primary">
                    Learn More
                  </Button>
                </CardActions>
              </Card>
            </Grid>
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
