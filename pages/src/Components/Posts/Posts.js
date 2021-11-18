import React from 'react';
import './Posts.css';

import Post from '../Post/Post';

class Posts extends React.Component {
  render() {
    return (
      <div className='Posts-Container'>
        {this.props?.posts.map((post, index) => (
          <Post post={post} key={index} />
        ))}
      </div>
    );
  }
}

export default Posts;
