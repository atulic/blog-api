import React from "react";
import "../styles/App.css";
import {BlogPost} from "./BlogPost";
import {CreateBlogPost} from "./CreateBlogPost";

const App: React.FC = () => {
    return (
        <>
            <CreateBlogPost/>
            <BlogPost/>
        </>);
};

export default App;
