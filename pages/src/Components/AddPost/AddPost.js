import React from 'react';
import './AddPost.css';
import { putPost } from '../../controllers/posts';

class AddPost extends React.Component {
  state = {
    username: '',
    title: '',
    type: 'select-type',
    content: '',
  };

  handleUsername = (e) => {
    const temp = e.target.value;
    this.setState({ username: temp });
  };

  handleTitle = (e) => {
    const temp = e.target.value;
    this.setState({ title: temp });
  };

  handleType = (e) => {
    const temp = e.target.value;
    this.setState({ type: temp, content: '' });
  };

  handleContent = (e) => {
    const temp = e.target.value;
    this.setState({ content: temp });
  };

  handleSubmit = (e) => {
    const post = { ...this.state };
    putPost(post).then((res) => {
      this.props.updateState(res);
      this.setState({
        username: '',
        title: '',
        type: 'select-type',
        content: '',
      });
    });
  };

  render() {
    const isDisabled = !(
      this.state.username &&
      this.state.title &&
      this.state.content
    );
    return (
      <div className='AddPost-Container'>
        <div className='AddPost-Username'>
          <input
            type='text'
            name='username'
            placeholder='Username'
            value={this.state.username}
            onChange={this.handleUsername}
          />
        </div>
        <div className='AddPost-Title'>
          <input
            type='text'
            name='Title'
            placeholder='Post title'
            value={this.state.title}
            onChange={this.handleTitle}
          />
        </div>
        <div className='AddPost-Type'>
          <select
            name='type'
            onChange={this.handleType}
            value={this.state.type}
          >
            <option value='select-type'>Post Type</option>
            <option value='text'>Text</option>
            <option value='media'>Media</option>
          </select>
        </div>
        {this.state.type === 'text' && (
          <div className='AddPost-Content'>
            <textarea
              type='text'
              name='content'
              placeholder='Text'
              rows='10'
              cols='40'
              value={this.state.content}
              onChange={this.handleContent}
            />
          </div>
        )}
        {this.state.type === 'media' && (
          <div className='AddPost-Media'>
            <input
              type='text'
              name='media'
              placeholder='Media URL'
              value={this.state.content}
              onChange={this.handleContent}
            />
          </div>
        )}
        <div className='AddPost-Submit'>
          <button disabled={isDisabled} onClick={this.handleSubmit}>
            Submit
          </button>
        </div>
      </div>
    );
  }
}

export default AddPost;
