import { render } from '@testing-library/svelte'
import Index from './App.svelte'

describe('index component', () => {
  test('should render component correctly', () => {
    const { container } = render(Index)

    //   expect(container).toContainHTML('<div></div>')
  })
})
