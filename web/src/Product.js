import React from 'react';

export default ({
  product,
  handleClick
}) => {

  const { name, image, description, price, stock } = product;

  return (
    <div className='product'>
      <img className='product-image' src={ image } />
      <h3>{ name }</h3>
      <span className="product-description">{ description }</span>
      <div className='product-figures'>
        <span><strong>Price: </strong>£{ price }</span>
        <span><strong>Stock: </strong>{ stock }</span>
      </div>
      <button
        className="product-btn"
        onClick={ _ => handleClick(product) }
      >
        Add to Cart
      </button>
    </div>
  )
}