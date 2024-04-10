// __tests__/App.test.js

import React from 'react';
import { render } from '@testing-library/react';
import MyApp from '../app/App';

describe('App', () => {
  it('renders the app component', () => {
    const { getByText } = render(<MyApp />);
    const linkElement = getByText(/Learn Next.js/i);
    expect(linkElement).toBeInTheDocument();
  });
});
