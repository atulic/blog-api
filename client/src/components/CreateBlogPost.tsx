import * as React from "react";
import { useState } from "react";
import Mutation from "react-apollo/Mutation";
import { POST_QUERY } from "../queries/fetchPostQuery";
import { POST_MUTATION } from "../mutations/createPostMutation";
import { CreatePost, CreatePost_create } from "../mutations/types/CreatePost";
import { FetchPosts } from "../queries/types/FetchPosts";
import { Button, Grid, TextField } from "@material-ui/core";

export const CreateBlogPost: React.FC = () => {
  const [post, setPost] = useState<Partial<CreatePost_create>>({
    title: "",
    content: ""
  });

  const textFieldStyle = { paddingLeft: "5px", paddingRight: "5px" };
  return (
    <Grid>
      <TextField
        value={post.title as string}
        margin="normal"
        onChange={e =>
          setPost({ title: e.target.value, content: post.content })
        }
        type="text"
        placeholder="Insert a post title"
        style={textFieldStyle}
      />
      <TextField
        value={post.content as string}
        margin="normal"
        onChange={e => setPost({ title: post.title, content: e.target.value })}
        type="text"
        placeholder="Some content for the post"
        style={textFieldStyle}
      />

      <Mutation<CreatePost, Partial<CreatePost_create>>
        mutation={POST_MUTATION}
        variables={{ title: post.title, content: post.content }}
        update={(store, { data }) => {
          if (!data || !data.create) {
            return;
          }

          const previous: FetchPosts | null = store.readQuery({
            query: POST_QUERY
          });
          previous && previous.posts && previous.posts.unshift(data.create);

          store.writeQuery({
            query: POST_QUERY,
            data: previous
          });
        }}
      >
        {postMutation => (
          <Button
            variant="contained"
            onClick={postMutation as any}
            style={{ margin: "10px" }}
          >
            Submit
          </Button>
        )}
      </Mutation>
    </Grid>
  );
};
