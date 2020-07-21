import routes from '../../../utils/routes'

interface SingleMenuBlock {
    title: string;
    url: string;
    children?: SingleMenuBlock[];
}

export default [
    {
        title: 'common.components.menu.home',
        url: routes.home
    },
    {
        title: 'common.components.menu.budget',
        url: routes.budget
    },
    {
        title: 'common.components.menu.shops',
        url: routes.shops
    },
    {
        title: 'common.components.menu.statistics',
        url: routes.stats
    },
    {
        title: 'common.components.menu.categorization',
        url: routes.categorization
    }
]
