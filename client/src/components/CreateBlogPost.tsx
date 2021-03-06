import * as React from "react";
import { useState } from "react";
import Mutation from "react-apollo/Mutation";
import { POST_QUERY } from "../queries/fetchPostQuery";
import { POST_MUTATION } from "../mutations/createPostMutation";
import { CreatePost, CreatePost_create } from "../mutations/types/CreatePost";
import { FetchPosts } from "../queries/types/FetchPosts";
import { Button, Grid, TextField } from "@material-ui/core";

export const CreateBlogPost: React.FC = () => {
  const [post, setPost] = useState<Omit<CreatePost_create, "__typename">>({
    title: "",
    content: ""
  });

  const textFieldStyle = {
    marginLeft: "5px",
    marginRight: "5px"
  };

  return (
    <Grid>
      <TextField
        value={post.title as string}
        margin="normal"
        onChange={e =>
          setPost({ title: e.target.value, content: post.content })
        }
        type="text"
        label="Post Title"
        placeholder="Insert a post title"
        style={textFieldStyle}
      />
      <TextField
        value={post.content as string}
        margin="normal"
        onChange={e => setPost({ title: post.title, content: e.target.value })}
        type="text"
        label="Post Content"
        placeholder="Insert some content"
        style={textFieldStyle}
      />

      <Mutation<CreatePost, Omit<CreatePost_create, "__typename">>
        mutation={POST_MUTATION}
        variables={{ title: post.title, content: post.content }}
        update={(store, { data }) => {
          if (!data || !data.create) {
            return;
          }

          const previous: FetchPosts | null = store.readQuery({
            query: POST_QUERY
          });

          previous?.posts?.unshift(data.create);

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
            style={{ margin: "20px 10px" }}
          >
            Submit
          </Button>
        )}
      </Mutation>
    </Grid>
  );
};
