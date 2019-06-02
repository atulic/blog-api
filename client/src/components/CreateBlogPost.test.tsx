import { mount, ReactWrapper } from "enzyme";
import { MockedProvider } from "react-apollo/test-utils";
import { CreateBlogPost } from "./CreateBlogPost";
import * as React from "react";
import { TextField } from "@material-ui/core";

describe("the CreateBlogPost component", () => {
  let wrapper: ReactWrapper;

  beforeEach(() => {
    wrapper = mount(
      <MockedProvider>
        <CreateBlogPost />
      </MockedProvider>
    );
  });

  it("renders inputs for the title and content", () => {
    const inputTitles = wrapper
      .find(TextField)
      .map(field => field.props().label);

    console.log(wrapper.debug());

    expect(inputTitles).toEqual(["Post Title", "Post Content"]);
  });
});
