import * as React from "react";
import { MockedProvider, MockedResponse } from "react-apollo/test-utils";
import { mount } from "enzyme";
import { POST_QUERY } from "../../queries/fetchPostQuery";
import { FetchPosts } from "../../queries/types/FetchPosts";
import { Loading } from "../Loading";
import { BlogPostList } from "./BlogPostList";

const mockPost: FetchPosts = {
  posts: [
    {
      title: "Title",
      content: "Content",
      __typename: "Post"
    }
  ]
};

const mocks: ReadonlyArray<MockedResponse> = [
  {
    request: {
      query: POST_QUERY
    },
    result: { data: mockPost }
  }
];

describe("The BlogPostList Component", () => {
  it("renders the loading state", async () => {
    const wrapper = mount(
      <MockedProvider mocks={mocks} addTypename={true}>
        <BlogPostList />
      </MockedProvider>
    ).childAt(0);

    expect(wrapper.find(Loading).exists()).toBe(true);
  });
});
