import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  createStyles,
  Grid,
  IconButton,
  makeStyles,
  Typography
} from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";
import React from "react";
import { FetchPosts_posts } from "../../queries/types/FetchPosts";

interface BlogCardProps {
  post: FetchPosts_posts;
}

const useStyles = makeStyles(
  createStyles({
    iconContainer: {
      display: "flex",
      justifyContent: "end",
      padding: "3px 3px 0px 0px"
    },

    content: {
      paddingTop: "0"
    }
  })
);

export const BlogCard: React.FC<BlogCardProps> = props => {
  const classes = useStyles();

  return (
    <Grid item xs={12} sm={6} md={4}>
      <Card>
        <Box className={classes.iconContainer}>
          <IconButton aria-label="Delete">
            <DeleteIcon fontSize="small" />
          </IconButton>
        </Box>
        <CardContent className={classes.content}>
          <Typography variant="h5">{props.post && props.post.title}</Typography>
          <Typography variant="body1">
            {props.post && props.post.content}
          </Typography>
        </CardContent>
        <CardActions>
          <Button size="small" color="primary">
            Learn More
          </Button>
        </CardActions>
      </Card>
    </Grid>
  );
};
