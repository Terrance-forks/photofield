import { LatestFetcher } from "./utils";
import { useTask } from "vue-concurrency";

export const host = "http://localhost:8080";

const fetchScene = LatestFetcher();

export async function get(endpoint, def) {
  const response = await fetch(host + endpoint);
  if (!response.ok) {
    if (def !== undefined) {
      return def;
    }
    console.error(response);
    throw new Error(response.statusText);
  }
  return await response.json();
}

export async function getRegions(x, y, w, h, sceneParams) {
  return get(`/regions?${sceneParams}&x=${x}&y=${y}&w=${w}&h=${h}`);
}

export async function getRegion(id, sceneParams) {
  return get(`/regions/${id}?${sceneParams}`);
}

export async function getCollections() {
  return get(`/collections`);
}

export async function getCollection(id) {
  return get(`/collections/` + id);
}

export function getTileUrl(level, x, y, tileSize, params) {
  let url = host + "/tiles";
  url += "?" + params;
  url += "&tileSize=" + tileSize;
  url += "&zoom=" + level;
  url += "&x=" + x;
  url += "&y=" + y;
  // for (const [key, value] of Object.entries(this.debug)) {
  //   url += "&debug" + key.slice(0, 1).toUpperCase() + key.slice(1) + "=" + (value ? "true" : "false");
  // }
  return url;
}

export async function getScene(params) {
  return get("/scenes?" + params);
}

export function useSceneTask() {
  return useTask(function*(_, params) {
    const scenes = yield get("/scenes?" + params);
    if (!scenes || scenes.length < 1) {
      throw new Error("Scene not found");
    }
    return scenes[0];
  });
}

export function useCollectionTask() {
  return useTask(function*(_, id) {
    return get("/collections/" + id);
  });
}

export function useRegionTask() {
  return useTask(function*(_, regionId, sceneParams) {
    if (sceneParams == null || regionId == null) {
      return null;
    }
    return getRegion(regionId, sceneParams);
  })
}



export default {
  get,
  getRegions,
  getCollections,
  getTileUrl,
}