import React from "react";
import { BlogPostList } from "./Cards/BlogPostList";
import { CreateBlogPost } from "./CreateBlogPost";
import "typeface-roboto";
import { Container } from "@material-ui/core";

const App: React.FC = () => {
  return (
    <Container>
      <CreateBlogPost />
      <BlogPostList />
    </Container>
  );
};

export default App;
