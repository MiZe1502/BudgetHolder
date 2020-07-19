<script lang="typescript">
    import {Category} from "../types";
    import EditCategoryActionElement
        from "../EditCategoryActionElement/EditCategoryActionElement.svelte";
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";

    export let data: Category = {};
</script>

<CommonActionElement>
    <EditCategoryActionElement data={data}/>
</CommonActionElement>

