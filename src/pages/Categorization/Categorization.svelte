<script lang="typescript">
    import {onMount} from "svelte";
    import {SideMainPadding} from "./style";
    import {mockData, mockCategories} from "./data";
    import CategoryRow from "./CategoryRow/CategoryRow.svelte";
    import CategoriesContainer
        from "./CategoriesContainer/CategoriesContainer.svelte";
    import {
        categoriesStatus,
        categories,
        simpleCategories,
        simpleCategoriesStatus,
        categoriesTotal,
        findParentCategory
    } from "../../stores/categories";
    import {LoadingStatus} from "../../stores/utils";

    onMount(async () => {
        categoriesStatus.set(LoadingStatus.Loading)
        simpleCategoriesStatus.set(LoadingStatus.Loading)

        setTimeout(() => {
            categoriesStatus.set(LoadingStatus.Finished)
            simpleCategoriesStatus.set(LoadingStatus.Finished)

            categories.set(mockData);
            simpleCategories.set(mockCategories);

            categoriesTotal.set(mockData.length);

            console.log(findParentCategory(17, $categories))
        }, 5000)
    })

</script>

<section>
    <CategoriesContainer categories={$categories} status={$categoriesStatus}/>
</section>
