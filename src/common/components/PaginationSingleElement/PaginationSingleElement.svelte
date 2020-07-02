<script lang="typescript">

    export let pageNumber: number = 0;
    export let onClick: (page: number) => {};
    export let isActive: boolean = false;
</script>

<style>
    .SingleButton {
        height: 25px;
        width: 25px;
        border: 1px solid #d6d2ce;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 4px;
        margin-right: 4px;
    }

    .SingleButton:hover {
        cursor: pointer;
        background-color: #9ef06e;
    }

    .ActiveButton {
        background-color: #5b9c35;
    }

</style>

<div on:click={() => onClick(pageNumber)} class="SingleButton" class:ActiveButton={isActive}>
    {pageNumber}
</div>