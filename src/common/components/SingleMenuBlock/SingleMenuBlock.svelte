<script lang="ts">
    import { Link } from "svelte-routing";

    export let title = "";
    export let url = "";
    export let isActive = false;
</script>

<style lang="scss">
    .SingleMenuBlock {
        height: 5vh;
        width: auto;
        background-color: #80cf47;
        padding-bottom: 8px;
        padding-top: 8px;
        padding-left: 12px;
        padding-right: 12px;
        min-width: 100px;
        justify-content: center;
    }

    .SingleMenuBlock:active,
    .SingleMenuBlock:hover {
        cursor: pointer;
        background-color: #5b9c35;
    }
</style>



<Link class="" to={url}>
    <div class="SingleMenuBlock Font422Black PrimaryBackground FlexHor">
        {title}
    </div>
</Link>
