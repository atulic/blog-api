import {gql} from "apollo-boost";
import * as React from "react";
import {useState} from "react";
import {Post} from "./types";
import Mutation from "react-apollo/Mutation";

const POST_MUTATION = gql`
    mutation ($title: String!, $content: String!) {
        create(title: $title, content: $content) {
      id
    }
  }`;

export const CreateBlogPost: React.FC = () => {
    const [post, setPost] = useState<Post>({title: "", content: ""});
    return (
        <div>
            <div>
                <input
                    value={post.title}
                    onChange={e => setPost({title: e.target.value, content: post.content})}
                    type="text"
                    placeholder="Insert a post title"
                />
                <input
                    className="mb2"
                    value={post.content}
                    onChange={e => setPost({title: post.title, content: e.target.value})}
                    type="text"
                    placeholder="Some content for the post"
                />
            </div>
            <Mutation<Post, Post> mutation={POST_MUTATION} variables={{title: post.title, content: post.content}}>
                {postMutation => <button onClick={postMutation as any}>Submit</button>}
            </Mutation>
        </div>
    );
};


