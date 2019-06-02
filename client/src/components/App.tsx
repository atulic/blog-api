import React from "react";
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
