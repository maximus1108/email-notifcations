import React from 'react';

export default ({
  name,
  quantity
}) => (
  <div className='cart-item'>
    <span>{ name }</span>
    <span>{ quantity }</span>
  </div>
)