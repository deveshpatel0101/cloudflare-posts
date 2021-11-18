import React from 'react';
import './Post.css';

class Post extends React.Component {
  render() {
    return (
      <div className='Post-Container'>
        <div className='Post-Username'>@{this.props.post.username}</div>
        <div className='Post-Title'>{this.props.post.title}</div>
        {this.props.post.type === 'text' && (
          <div className='Post-Content'>{this.props.post.content}</div>
        )}
        {this.props.post.type === 'media' && (
          <div className='Post-Content'>
            <img src={this.props.post.content} alt={this.props.post.title} />
          </div>
        )}
      </div>
    );
  }
}

export default Post;
