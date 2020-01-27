import React, { useState, useCallback } from 'react';

export default ({ onCancel, onSubmit }) => {

  const [ user, setUser ] = useState({
    firstName: '',
    surname: '',
    email: '',
  });

  const setProp = useCallback(prop => e => setUser({
    ...user,
    [prop]: e.target.value
  }), [user]);

  const handleSubmit = useCallback(e => {
    e.preventDefault() 
    onSubmit(user);
  }, [user]);

  console.log(user)

  return (
    <div className="modal-bg">
      <div className="modal-content">
        <form  onSubmit={ handleSubmit } >
          <div>
            <label>First name:</label>
            <input type="text" required onChange={ setProp('firstName') }/>
          </div>
          <div>
            <label>Surname:</label>
            <input type="text" required onChange={ setProp('surname') }/>
          </div>
          <div>
            <label>Email:</label>
            <input type="email" required onChange={ setProp('email') }/>
          </div>
          <div>
            <button>Buy</button>
            <button onClick={ onCancel }>Cancel</button>
          </div>
        </form>
      </div>
    </div>
  )
}