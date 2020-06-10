import routes from '../../utils/routes'

interface SingleMenuBlock {
    title: string;
    url: string;
    children?: SingleMenuBlock[];
}

export default [
  {
    title: 'Home',
    url: routes.home
  },
  {
    title: 'Budget',
    url: routes.budget
  },
  {
    title: 'Shops',
    url: routes.shops
  },
  {
    title: 'Statistics',
    url: routes.stats
  }
]
