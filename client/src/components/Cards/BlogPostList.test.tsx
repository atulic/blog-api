import * as React from "react";
import { MockedProvider, MockedResponse } from "react-apollo/test-utils";
import { mount, ReactWrapper } from "enzyme";
import { POST_QUERY } from "../../queries/fetchPostQuery";
import { FetchPosts } from "../../queries/types/FetchPosts";
import { Loading } from "../Loading";
import { BlogPostList } from "./BlogPostList";
import { BlogCard } from "./BlogCard";

let wrapper: ReactWrapper;

const mockPosts: FetchPosts = {
  posts: [
    {
      content: "Content",
      title: "Title",
      __typename: "Post"
    }
  ]
};

const mocks: ReadonlyArray<MockedResponse> = [
  {
    request: {
      query: POST_QUERY
    },
    result: { data: mockPosts }
  }
];

describe("The BlogPostList Component", () => {
  beforeEach(() => {
    jest.useFakeTimers();
    wrapper = mount(
      <MockedProvider mocks={mocks} addTypename={true}>
        <BlogPostList />
      </MockedProvider>
    );
  });

  it("renders the loading state", async () => {
    expect(wrapper.find(Loading).exists()).toBe(true);
  });

  it("renders the Blog Card and passes the result of the query", async () => {
    jest.runOnlyPendingTimers();
    wrapper.update();

    expect(wrapper.find(BlogCard).props().post).toEqual(mockPosts!.posts![0]);
  });
});
