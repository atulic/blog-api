import React from "react";
import { BlogPost } from "./BlogPost";
import { CreateBlogPost } from "./CreateBlogPost";
import "typeface-roboto";
import { Container } from "@material-ui/core";

const App: React.FC = () => {
  return (
    <Container>
      <CreateBlogPost />
      <BlogPost />
    </Container>
  );
};

export default App;
