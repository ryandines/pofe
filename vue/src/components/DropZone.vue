<template>
  <div>
    <sp-h3>Calculate Proof</sp-h3>
    <div id="drop-zone" @drop.prevent="hash" @dragover.prevent>
      {{ hashed || placeholderText }}
    </div>
    <sp-button
      @click="submit"
      :loading="flight"
      :disabled="hashed == '' || flight"
      >Submit Proof as a new Claim</sp-button
    >
  </div>
</template>

<style scoped>
#drop-zone {
  font-family: "Inter", "Helvetica", sans-serif;
  height: 20rem;
  line-height: 20rem;
  font-size: 1.25rem;
  background-color: rgb(247, 247, 247);
  text-align: center;
  margin-bottom: 2rem;
  border-radius: 0.5rem;
}
</style>
