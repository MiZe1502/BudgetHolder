<script lang="typescript">
    import {Category} from "../types";
    import EditCategoryActionElement
        from "../EditCategoryActionElement/EditCategoryActionElement.svelte";
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import RemoveCategoryActionElement
        from "../RemoveCategoryActionElement/RemoveCategoryActionElement.svelte";

    export let data: Category = {};
</script>

<CommonActionElement>
    <EditCategoryActionElement data={data}/>
    <RemoveCategoryActionElement data={data}/>
</CommonActionElement>

